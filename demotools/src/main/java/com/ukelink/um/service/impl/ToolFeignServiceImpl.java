package com.ukelink.um.service.impl;

import com.ukelink.um.feign.ToolFeignClient;
import com.ukelink.um.service.ToolFeignService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @ClassName ToolFeignServiceImpl
 * @Description TODO
 * @Author chuang.gao
 * @Date 2019-7-30 17:55
 * @Version 1.0
 */
@Service
public class ToolFeignServiceImpl implements ToolFeignService {
    @Autowired
    private ToolFeignClient toolFeignClient;
    @Override
    public String testFeignServer()
    {
        System.out.println("testFeignServer function");
        return toolFeignClient.testFeign(12345);
    }
}
