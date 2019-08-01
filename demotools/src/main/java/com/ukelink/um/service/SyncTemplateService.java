package com.ukelink.um.service;

import com.ukelink.um.proto.Message;
import com.ukelink.um.proto.login.Accountinfo;
import com.ukelink.um.proto.login.Login;
import okhttp3.*;
import okio.BufferedSink;
import org.springframework.stereotype.Service;

import java.io.ByteArrayOutputStream;
import java.io.IOException;

import static java.nio.charset.StandardCharsets.UTF_8;

@Service
public class SyncTemplateService {

    private static final RequestBody requestBody1;
    private static Request request;
    private final OkHttpClient build = new OkHttpClient().newBuilder().build();

    static {
        requestBody1=new RequestBody() {
            Message.SyncMessageRequest syncMessageRequest;

            @Override
            public MediaType contentType()
            {
                return MediaType.parse("application/protobuf");
            }
            @Override
            public long contentLength() throws IOException {
                Message.SyncMessageRequest.Builder syncRequest = Message.SyncMessageRequest.newBuilder();
                syncRequest.setUserId("112233");
                syncRequest.setClientID("112233");
//                Message.SyncMessageRequest.SyncKey.Builder synckey = Message.SyncMessageRequest.SyncKey.newBuilder();
//                synckey.setKey(1);
//                synckey.setWinSize(10);
//                synckey.setMailId("101");
//                syncRequest.addSyncKey(synckey);
                syncMessageRequest = syncRequest.build();
                return 4 + 8 + 12 + 8L+ syncMessageRequest.getSerializedSize();
            }
            @Override
            public void writeTo(BufferedSink buffersink) throws IOException {
                ByteArrayOutputStream loginStream = new ByteArrayOutputStream();
                syncMessageRequest.writeTo(loginStream);
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
                    .url("http://10.100.93.47:8802/um/mail/newsync")
                    .post(requestBody1)
                    .removeHeader("Content-Encoding")
                    .addHeader("Content-Length", String.valueOf(requestBody1.contentLength()))
                    .build();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public String sync() {
//        try {
//            Response response = build.newCall(request).execute();
//            System.out.println("sync response: " + response.toString());
//            return response.body().string();
            return "0";
            //response.close();
//        } catch (IOException e) {
//            e.printStackTrace();
//        }

//        return "ok";
    }
}
