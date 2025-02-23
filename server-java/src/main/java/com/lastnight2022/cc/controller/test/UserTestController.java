package com.lastnight2022.cc.controller.test;

import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class UserTestController {
    @GetMapping("/public")
    public String publicApi() {
        return "公开访问接口";
    }

    @GetMapping("/private")
    @PreAuthorize("hasAnyRole('USER', 'ADMIN')")
    public String privateApi() {
        return "需要登录的接口";
    }

    @GetMapping("/admin-only")
    @PreAuthorize("hasRole('ADMIN')")
    public String adminApi() {
        return "管理员专属接口";
    }
}
