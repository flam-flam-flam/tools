package com.ukelink.um.controller;

import cn.hutool.core.codec.Base64Decoder;
import com.google.protobuf.InvalidProtocolBufferException;
import com.ukelink.um.proto.Message;
import com.ukelink.um.redis.RedisDAO;
import com.ukelink.um.service.SendTemplateService;
import com.ukelink.um.service.ToolFeignService;
import com.ukelink.um.util.Md5Util;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.security.NoSuchAlgorithmException;
import java.util.ArrayList;
import java.util.List;

import static com.ukelink.um.redis.LuaScript.TESTLUA;

@RestController
@Slf4j
public class SendMessage {
    @Autowired
    private RedisDAO redisDAO;
//    @Autowired
//    private RestTemplateService restTemplateService;
    @Autowired
    private SendTemplateService sendTemplateService;
    @Autowired
    private ToolFeignService toolFeignService;
    @RequestMapping("/um/message/send")
    public String index() {

        return sendTemplateService.sync();
    }

    @RequestMapping("/um/message/send2")
    public String index2(@RequestParam("id") Integer id) throws InvalidProtocolBufferException {
         // String byte2 = "ChoKCDAwMDAwMDAwEgdzdWNjZXNzKMqjwObDLRJvCAoaawppEg4IBBIKEICAgKzjloSFEAobCAESBhCbh5CcAxoPODYyNjU4MDMwMDk5MzYyGAcqGDVkM2VhYzBjZDVkZWYwNDU0ZDg2ZTVmMTi+pcDmwy1CFwoVCKW0w4kDEg04NjE3NzA5MjIxNzIy";
//        String byte2 = "CmkKGwgBEgYQm4eQnAMaDzg2MjY1ODAzMDA5OTM2MhIOCAQSChCAgICqwuWBhRAYByoYNWQzZWEyM2VkNWRlZjA0NTRkODZlNWU4OP6Kp+XDLUIXChUIpbTDiQMSDTg2MTc3MDkyMjE3MjI=";
//        String byte2= "CnkKIAgBEgsSCTg2NDI4OTY5MRoPODYyNjU4MDMwMDk5MzYyEhkIBBIVEhMxMTU1Njg1ODczOTQ2MTk4MDE2GAcqGDVkM2U2YzA1ZDVkZTdiYTFiY2I3ZGEyZjiT69fewy1CFwoVCKW0w4kDEg04NjE3NzA5MjIxNzIy";
//        Message.Data.Builder data1 = Message.Data.newBuilder();
//        data1.setContent(ByteString.copyFromUtf8("I-am-content"));
//        data1.setAction(Message.Data.MessageType.MESSAGE_ACTION_ADD_GROUP_MEMBER);
//        data1.setIndexId("11111");
//        data1.build();
        //byte[] bs = Base64Decoder.decode(byte2);
        byte[] bs = {10, 105, 10, 27, 8, 1, 18, 6, 16, -101, -121, -112, -100, 3, 26, 15, 56, 54, 50, 54, 53, 56, 48, 51, 48, 48, 57, 57, 51, 54, 50, 18, 14, 8, 4, 18, 10, 16, -128, -128, -128, -44, -75, -116, -122, -123, 16, 24, 7, 42, 24, 53, 100, 51, 101, 98, 51, 101, 51, 100, 53, 100, 101, 102, 48, 52, 53, 52, 100, 56, 54, 101, 53, 102, 55, 56, -113, -29, -70, -25, -61, 45, 66, 23, 10, 21, 8, -91, -76, -61, -119, 3, 18, 13, 56, 54, 49, 55, 55, 48, 57, 50, 50, 49, 55, 50, 50};
        Message.SendMessageRequest sendMessageRequest = Message.SendMessageRequest.parseFrom(bs);
        //Message.SyncMessageResponse sendMessageRequest = Message.SyncMessageResponse.parseFrom(bs);
        sendMessageRequest.toBuilder();
        Message.Data data = sendMessageRequest.getData();
        System.out.println("id:" + id);
        System.out.println("send response11111w2223334444: " + sendMessageRequest.toString());
        //Message.Data data = Message.Data.parseFrom(bs);
        return sendMessageRequest.toString();
    }
    @RequestMapping("/um/message/send3")
    public String index3() throws InvalidProtocolBufferException, NoSuchAlgorithmException {
        return Md5Util.md5("456","456");
    }
    @RequestMapping("/um/get/md5")
    public String getMd5(@RequestParam("username") String username, @RequestParam("password") String password) throws NoSuchAlgorithmException {
        System.out.println(username + password);
        String md5Value = Md5Util.md5(username,password);
        String key = "turn/realm/mytest/user/" + username + "/key";
        System.out.println("key:" + key);
        redisDAO.set(null, key, md5Value);
        String feignTest = toolFeignService.testFeignServer();
        System.out.println("feign:" + feignTest);
        String Result = md5Value + "==========\r\n============" + feignTest;
        List<String> keys = new ArrayList<>();
        keys.add("luatest1");
        List<String> argv = new ArrayList<>();
        argv.add("2345");
        String value = (String)redisDAO.luaEvalSha1(null,TESTLUA,"lua22222",keys,argv);

//        String value = (String)redisDAO.luaEval(null,TESTLUA,null,keys,argv);
        System.out.println("value:" + value);
        log.info("[lid:{}] createGroup() requestBody: {}", value, argv.toString());
        return Result;
    }
}


