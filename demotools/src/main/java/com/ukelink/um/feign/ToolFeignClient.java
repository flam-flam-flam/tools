package com.ukelink.um.feign;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;

/**
 * @InterfaceName ToolFeignClient
 * @Description TODO
 * @Author chuang.gao
 * @Date 2019-7-30 17:11
 * @Version 1.0
 */
@FeignClient("UM-TOOL")
public interface ToolFeignClient {
    @GetMapping("/um/message/send2")
    String testFeign(@RequestParam("id") Integer id);
}
