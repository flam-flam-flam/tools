package com.ukelink.um.controller;

import com.ukelink.um.service.SyncTemplateService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SyncMessage {
    @Autowired
    private SyncTemplateService syncTemplateService;
    @RequestMapping("/um/mail/newsync")
    public String index(){
        return syncTemplateService.sync();
    }

}
