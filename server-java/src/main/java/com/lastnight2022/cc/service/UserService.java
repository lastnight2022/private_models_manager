package com.lastnight2022.cc.service;

import com.lastnight2022.cc.mappers.UserMapper;
import com.lastnight2022.cc.models.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class UserService {
    @Autowired
    private final UserMapper userMapper;
    private final PasswordEncoder passwordEncoder;

    public UserService(UserMapper userMapper, PasswordEncoder passwordEncoder) {
        this.userMapper = userMapper;
        this.passwordEncoder = passwordEncoder;
    }

    // 注册用户
    public void registerUser(User user) {
        // 密码加密
        String encodedPassword = passwordEncoder.encode(user.getPasswordHash());
        user.setPasswordHash(encodedPassword);
        // 设置默认角色
        user.setRole(User.Role.user);

        userMapper.insertUser(user);
    }

    // 根据用户名查询
    public User findByUsername(String username) {
        return userMapper.selectUserByUsername(username);
    }
}