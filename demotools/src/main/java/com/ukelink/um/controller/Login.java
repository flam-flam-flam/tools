package com.ukelink.um.controller;

import com.ukelink.um.service.LoginTemplateService;
import com.ukelink.um.service.RestTemplateService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class Login {
    @Autowired
    private LoginTemplateService loginTemplateService;

    @RequestMapping("/um/account/login")
    public String index() {
        return loginTemplateService.login();
    }
}
