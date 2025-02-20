package com.lastnight2022.cc.mapper;

import com.lastnight2022.cc.models.User;
import org.apache.ibatis.annotations.Mapper;

import java.util.List;

@Mapper
public interface UserMapper {
    // 根据ID查询用户
    User selectUserById(Long id);

    // 插入用户（返回自增ID）
    int insertUser(User user);

    // 根据用户名查询用户
    User selectUserByUsername(String username);

    // 查询所有用户
    List<User> selectAllUsers();
}
