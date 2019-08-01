package com.ukelink.um.service;

import com.google.protobuf.ByteString;
import com.ukelink.um.bean.VmpHeader;
import com.ukelink.um.proto.Clientinfo;
import com.ukelink.um.proto.Message;
import com.ukelink.um.proto.Vmmbody;
import com.ukelink.um.proto.Vmmheader;
import okhttp3.*;
import okio.*;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.io.ByteArrayOutputStream;
import java.io.DataInputStream;
import java.io.IOException;

import static java.nio.charset.StandardCharsets.UTF_8;

@Service
public class RestTemplateService {

    private static final RequestBody requestBody1;
    private static Request request;
    private final OkHttpClient build = new OkHttpClient().newBuilder().build();


    //@Autowired
    //RestTemplate restTemplate;

    static {
        requestBody1=new RequestBody() {
            Message.SendMessageRequest request;

            @Override
            public MediaType contentType()
            {
                return MediaType.parse("application/protobuf");
            }
            @Override
            public long contentLength() throws IOException {
                Clientinfo.ClientInfo.UserValue.Builder userValue = Clientinfo.ClientInfo.UserValue.newBuilder();
               // userValue.setId("112233");
                userValue.setNumber("138");
                Clientinfo.ClientInfo.Builder fromClientInfo = Clientinfo.ClientInfo.newBuilder();
                fromClientInfo.setClientId("aaa");
                fromClientInfo.setDeviceType("PC");
                fromClientInfo.setUserType(Clientinfo.ClientInfo.UserType.sim);
                fromClientInfo.setUserValue(userValue);

                Vmmheader.VmmHeader.Builder vmmheader = Vmmheader.VmmHeader.newBuilder();
                vmmheader.setFromUser(fromClientInfo);
                vmmheader.addToUser(fromClientInfo);
                vmmheader.setConversationType(1);
                vmmheader.setAction(1);
                vmmheader.setLanguage("zh");
                vmmheader.setOriginalMsgId("I-am-msgid");
//                vmmheader.setActionMsgId("aaa");

                Vmmbody.VmmBody.Builder vmmbody = Vmmbody.VmmBody.newBuilder();
                vmmbody.setContentType(1);
                vmmbody.setMediaType(1);
                vmmbody.setContent(ByteString.copyFromUtf8("I-am-content"));
                Message.SendMessageRequest.Builder vmm = Message.SendMessageRequest.newBuilder();
               // vmm.setVmmHeader(vmmheader);
                //vmm.setVmmBody(vmmbody);
                request = vmm.build();
                return 4 + 8 + 12 + 8L+request.getSerializedSize();
            }
            @Override
            public void writeTo(BufferedSink sink) throws IOException {
                //OutputStream outputStream=sink.outputStream();
                ByteArrayOutputStream vmmStream = new ByteArrayOutputStream();
                request.writeTo(vmmStream);
//                ByteArrayOutputStream byteArrayOutputStream = new ByteArrayOutputStream();
//                DataOutputStream dis = new DataOutputStream(byteArrayOutputStream);
                String token = "12345678";
                sink.writeByte(0xEF);
                sink.writeByte(0);
                sink.writeByte(0);
                sink.writeByte(token.length());
                sink.write(token.getBytes(UTF_8));
                sink.writeShort(10);
                sink.writeLong(112233);
                sink.writeShort(11);
                byte[] bytes = vmmStream.toByteArray();
                sink.writeInt(bytes.length);
                sink.writeInt(bytes.length);
                sink.write(bytes);
            }
        };
    }

    static {
        try {
            request = new Request.Builder()
                    .url("http://10.100.93.47:8802/um/message/send")
                    .post(requestBody1)
                    .removeHeader("Content-Encoding")
                    .addHeader("Content-Length", String.valueOf(requestBody1.contentLength()))
                    .build();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public String sendMessage() {

//        OkHttpClient build = new OkHttpClient().newBuilder().build();

       try {
           Response response = build.newCall(request).execute();
           return response.body().string();
           // response.close();
//            build.newCall(request).execute();
//            if (response.code() == 200){
//
//            }
//            Source source = Okio.source(response.body().byteStream());
//            BufferedSource dis = Okio.buffer(source);
//
////            DataInputStream dis = new DataInputStream();
////            System.out.println("magic: " + response.toString());
////            System.out.println("response.code: " + response.toString());
//            VmpHeader vmpHeader = new VmpHeader();
//            vmpHeader.setMagic(dis.readByte());
//            vmpHeader.setHeaderLen((byte) (dis.readByte() >>> 2));
//            byte zipFlagAndEncrypType = dis.readByte();
//            vmpHeader.setEncrypType((byte) (zipFlagAndEncrypType >>> 2 & 63));
//            vmpHeader.setZipFlag((byte) (zipFlagAndEncrypType & 3));
//            vmpHeader.setTokenLen(dis.readByte());
//            byte[] tokenArray = new byte[vmpHeader.getTokenLen()];
//            dis.read(tokenArray);
//            vmpHeader.setToken(new String(tokenArray));
//            vmpHeader.setVersion(dis.readShort());
//            vmpHeader.setUserId(dis.readLong());
//            vmpHeader.setCgiId(dis.readShort());
//            vmpHeader.setLenOrgBody(dis.readInt());
//            vmpHeader.setLenCompressed(dis.readInt());
//            byte[] bodyArray = new byte[vmpHeader.getLenCompressed()];
//            dis.read(bodyArray);
//            Message.SendMessageResponse sendMessageResponse = Message.SendMessageResponse.parseFrom(bodyArray);
        } catch (IOException e) {
            e.printStackTrace();
        }

        return "ok";
    }

}
