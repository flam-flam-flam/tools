package com.ukelink.um.service;

import com.ukelink.um.bean.MySendRequestBody;
import com.ukelink.um.util.IDUtil;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import org.springframework.stereotype.Service;

import java.io.IOException;

@Service
public class SendTemplateService {
    private static final OkHttpClient build = new OkHttpClient().newBuilder().build();
    public String sync() {
        MySendRequestBody requestBody=new MySendRequestBody(IDUtil.getId());
        Request httprequest = new Request.Builder()
                .url("http://10.100.93.47:8802/um/message/send")
                .post(requestBody)
                .removeHeader("Content-Encoding")
                .build();
        try {
            Response response = build.newCall(httprequest).execute();
            System.out.println("send response11111w222: " + response.toString());
            return response.body().string();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return "ok";
    }
}
