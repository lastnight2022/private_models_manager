# create database AI_Models_Manager;

use AI_Models_Manager;

-- 模型元数据表（models）
--  CREATE TABLE models (
--                          id SERIAL PRIMARY KEY,
--                          name VARCHAR(255) UNIQUE NOT NULL,  -- 模型名称（唯一标识）
--                          description TEXT,
--                          created_by INT REFERENCES users(id),
--                          created_at TIMESTAMP DEFAULT NOW(),
--                          is_public BOOLEAN DEFAULT false     -- 是否公开可见
-- );

-- 模型版本表（model_versions）
CREATE TABLE model_versions (
                                id SERIAL PRIMARY KEY,
                                model_id INT REFERENCES models(id),
                                version VARCHAR(50) NOT NULL,       -- 语义化版本（如v1.0.2）
                                file_path VARCHAR(512) NOT NULL,    -- 模型文件存储路径
                                config JSON not null ,              -- Ollama Modelfile 配置
                                status ENUM('uploading', 'active', 'disabled') DEFAULT 'uploading',
                                created_at TIMESTAMP DEFAULT NOW()
);

-- 模型权限表（model_permissions）
CREATE TABLE model_permissions (
                                   model_id INT REFERENCES models(id),
                                   user_id INT REFERENCES users(id),
                                   access_level ENUM('read', 'write', 'admin'),
                                   PRIMARY KEY (model_id, user_id)
);

-- 模型市场元数据表（market_models）
CREATE TABLE market_models (
                               id SERIAL PRIMARY KEY,
                               source VARCHAR(20) NOT NULL,  -- 来源（huggingface/ollama/custom）
                               source_id VARCHAR(255) NOT NULL, -- 源站唯一ID（如HF的model_id）
                               name VARCHAR(255) NOT NULL,
                               description TEXT,
                               author VARCHAR(100),
                               task_type VARCHAR(50),        -- 任务类型（text-classification等）
                               framework VARCHAR(50),        -- 框架（PyTorch/GGUF等）
                               downloads INT DEFAULT 0,
                               rating FLOAT DEFAULT 0.0,
                               last_synced TIMESTAMP,        -- 最后同步时间
                               UNIQUE(source, source_id)     -- 防止重复记录
);

-- 用户交互表（user_model_actions）
CREATE TABLE user_model_actions (
                                    user_id INT REFERENCES users(id),
                                    model_id INT REFERENCES market_models(id),
                                    action_type VARCHAR(20),      -- download/star/rate
                                    rating INT CHECK (rating BETWEEN 1 AND 5),
                                    comment TEXT,
                                    created_at TIMESTAMP DEFAULT NOW(),
                                    PRIMARY KEY (user_id, model_id, action_type)
);

-- 镜像源配置表（mirror_sources）
CREATE TABLE mirror_sources (
                                id SERIAL PRIMARY KEY,
                                name VARCHAR(100) UNIQUE,     -- 镜像源名称（如"HuggingFace中国站"）
                                endpoint VARCHAR(255),        -- API地址（如https://hf-mirror.com）
                                type VARCHAR(20) NOT NULL,    -- huggingface/ollama/custom
                                is_public BOOLEAN DEFAULT true,
                                priority INT DEFAULT 0        -- 优先级（数值越高优先使用）
);

-- 资源监控表（时序数据）
CREATE TABLE monitoring_metrics (
                                    time TIMESTAMP NOT NULL DEFAULT NOW(),
                                    instance_id VARCHAR(50),      -- 实例标识（用于集群部署）
                                    metric_name VARCHAR(50),      -- 如 cpu_usage、gpu_mem_used
                                    metric_value DOUBLE PRECISION,
                                    labels JSON                  -- 扩展标签（如gpu_index:0）
);
CREATE INDEX idx_metrics_time ON monitoring_metrics (time);

-- 请求性能表
CREATE TABLE request_stats (
                               request_id char(36) PRIMARY KEY DEFAULT (UUID()),
                               user_id INT REFERENCES users(id),
                               model_name VARCHAR(255),
                               prompt_length INT,
                               total_tokens INT,
                               start_time TIMESTAMP,
                               end_time TIMESTAMP,
                               tokens_per_sec FLOAT
);

CREATE TABLE alert_rules (
                             id SERIAL PRIMARY KEY,
                             metric_name VARCHAR(50) NOT NULL,
                             `condition` VARCHAR(20) NOT NULL,  -- '>', '<', '=='
                             threshold FLOAT NOT NULL,
                             duration time NOT NULL,      -- 持续时长（如5分钟）
                             severity VARCHAR(10)             -- critical/warning
);


-- ----------------------------
-- 用户表 (users)
-- ----------------------------
CREATE TABLE `users` (
                         `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户唯一ID',
                         `username` VARCHAR(50) NOT NULL COMMENT '用户名（唯一）',
                         `password_hash` VARCHAR(255) NOT NULL COMMENT '密码哈希值（加盐存储）',
                         `email` VARCHAR(100) NOT NULL COMMENT '邮箱（唯一）',
                         `role` ENUM('admin', 'user') NOT NULL DEFAULT 'user' COMMENT '用户角色',
                         `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
                         `last_login` DATETIME COMMENT '最后登录时间',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `uniq_username` (`username`),
                         UNIQUE KEY `uniq_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- 对话记录表 (conversations)
-- ----------------------------
CREATE TABLE `conversations` (
                                 `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '对话ID',
                                 `user_id` INT UNSIGNED NOT NULL COMMENT '关联用户ID',
                                 `model_name` VARCHAR(50) NOT NULL COMMENT '使用的模型名称（如llama2）',
                                 `input_text` TEXT NOT NULL COMMENT '用户输入内容',
                                 `output_text` TEXT NOT NULL COMMENT '模型生成内容',
                                 `timestamp` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '对话时间',
                                 PRIMARY KEY (`id`),
                                 KEY `idx_user_id` (`user_id`),
                                 KEY `idx_timestamp` (`timestamp`),
                                 CONSTRAINT `fk_conversation_user`
                                     FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
                                         ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- 模型管理表 (models)
-- ----------------------------
CREATE TABLE `models` (
                          `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '模型ID',
                          `name` VARCHAR(50) NOT NULL COMMENT '模型名称（唯一）',
                          `config_path` VARCHAR(255) NOT NULL COMMENT 'Ollama模型配置文件路径',
                          `created_by` INT UNSIGNED NOT NULL COMMENT '创建者（用户ID）',
                          `is_public` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否公开（1-是，0-否）',
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `uniq_name` (`name`),
                          CONSTRAINT `fk_model_creator`
                              FOREIGN KEY (`created_by`) REFERENCES `users` (`id`)
                                  ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- 权限表 (permissions)
-- ----------------------------
CREATE TABLE `permissions` (
                               `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
                               `model_id` INT UNSIGNED NOT NULL COMMENT '模型ID',
                               `access_level` ENUM('read', 'write', 'admin') NOT NULL DEFAULT 'read' COMMENT '权限级别',
                               PRIMARY KEY (`user_id`, `model_id`),
                               CONSTRAINT `fk_permission_user`
                                   FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
                                       ON DELETE CASCADE,
                               CONSTRAINT `fk_permission_model`
                                   FOREIGN KEY (`model_id`) REFERENCES `models` (`id`)
                                       ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- 操作日志表 (audit_logs)
-- ----------------------------
CREATE TABLE `audit_logs` (
                              `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志ID',
                              `user_id` INT UNSIGNED NOT NULL COMMENT '操作用户ID',
                              `action` VARCHAR(50) NOT NULL COMMENT '操作类型（login/delete等）',
                              `target_id` INT UNSIGNED COMMENT '操作目标ID（如对话ID）',
                              `timestamp` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_user_id` (`user_id`),
                              KEY `idx_timestamp` (`timestamp`),
                              CONSTRAINT `fk_audit_user`
                                  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
                                      ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;