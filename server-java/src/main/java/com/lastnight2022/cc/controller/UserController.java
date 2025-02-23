package com.lastnight2022.cc.controller;

import com.lastnight2022.cc.models.User;
import com.lastnight2022.cc.service.UserService;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import io.swagger.annotations.ApiParam;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;

@RestController
@RequestMapping("/")
@Api(tags = "用户管理接口", description = "用户注册、查询等操作") // 模块分类标签[10,11](@ref)
public class UserController {
    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    @PostMapping("/register")
    @ApiOperation(value = "用户注册", notes = "密码自动加密，默认角色为普通用户")
    public ResponseEntity<Void> registerUser(
            @ApiParam(value = "用户信息（需包含用户名和密码）", required = true)
            @RequestBody @Valid User user) {
        userService.registerUser(user);
        return ResponseEntity.status(HttpStatus.CREATED).build();
    }

    @GetMapping("/users/{username}")
    @ApiOperation(value = "查询用户信息", notes = "根据用户名获取用户详情")
    public ResponseEntity<User> findByUsername(
            @ApiParam(value = "用户名", example = "john_doe", required = true)
            @PathVariable String username) {
        User user = userService.findByUsername(username);
        return user != null ?
                ResponseEntity.ok(user) :
                ResponseEntity.notFound().build();
    }
}