package com.ukelink.um.service;

import com.ukelink.um.proto.login.Accountinfo;
import com.ukelink.um.proto.login.Deviceinfo;
import com.ukelink.um.proto.login.Login;
import okhttp3.*;
import okio.BufferedSink;
import org.springframework.stereotype.Service;

import java.io.ByteArrayOutputStream;
import java.io.IOException;

import static java.nio.charset.StandardCharsets.UTF_8;
@Service
public class LoginTemplateService {

    private static final RequestBody requestBody1;
    private static Request request;
    private final OkHttpClient build = new OkHttpClient().newBuilder().build();

    static {
        requestBody1=new RequestBody() {
            Login.LoginRequest request;

            @Override
            public MediaType contentType()
            {
                return MediaType.parse("application/protobuf");
            }
            @Override
            public long contentLength() throws IOException {
                Accountinfo.AccountInfo.Builder accountInfo = Accountinfo.AccountInfo.newBuilder();
                Deviceinfo.DeviceInfo.Builder deviceInfo = Deviceinfo.DeviceInfo.newBuilder();
                Login.LoginRequest.Builder loginRequst = Login.LoginRequest.newBuilder();
                accountInfo.setUserCode("11111111");
                accountInfo.setCountryCode("86");
                accountInfo.setMvnoCode("1");
                accountInfo.setVerifyCode("111");
                accountInfo.setPartnerCode("222");
                accountInfo.setPartnerCode("333");
                accountInfo.setPassword("444");
                deviceInfo.setDeviceId("11111111");
                deviceInfo.setDeviceType("mobile");
                deviceInfo.setHardwareVersion("1111");
                deviceInfo.setMobileBrand("2");
                deviceInfo.setSoftVersion("3");
                deviceInfo.setSystemVersion("4");
                deviceInfo.setImsi("5");
                deviceInfo.setImei("6");
                deviceInfo.setLangType("zh");
                deviceInfo.setPushPlatform("huawei");
                deviceInfo.setPushtoken("7");
                deviceInfo.setPushtokenType("8");

                loginRequst.setLoginType("phone");
                loginRequst.setAutoLogin(true);
                loginRequst.setAccountInfo(accountInfo);
                loginRequst.setDeviceInfo(deviceInfo);

                request = loginRequst.build();
                return 4 + 8 + 12 + 8L+ request.getSerializedSize();
            }
            @Override
            public void writeTo(BufferedSink buffersink) throws IOException {
                ByteArrayOutputStream loginStream = new ByteArrayOutputStream();
                request.writeTo(loginStream);
                String token = "12345678";
                buffersink.writeByte(0xEF);
                buffersink.writeByte(0);
                buffersink.writeByte(0);
                buffersink.writeByte(token.length());
                buffersink.write(token.getBytes(UTF_8));
                buffersink.writeShort(10);
                buffersink.writeLong(112233);
                buffersink.writeShort(11);
                byte[] bytes = loginStream.toByteArray();
                buffersink.writeInt(bytes.length);
                buffersink.writeInt(bytes.length);
                buffersink.write(bytes);
            }
        };
    }

    static {
        try {
            request = new Request.Builder()
                    .url("http://10.100.93.47:8802/um/account/login")
                    .post(requestBody1)
                    .removeHeader("Content-Encoding")
                    .addHeader("Content-Length", String.valueOf(requestBody1.contentLength()))
                    .build();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public String login() {
        try {
            Response response = build.newCall(request).execute();
            return response.body().string();
            //response.close();
        } catch (IOException e) {
            e.printStackTrace();
        }

        return "ok";
    }
}
