package com.ukelink.um.bean;

import com.google.protobuf.ByteString;
import com.ukelink.um.proto.Clientinfo;
import com.ukelink.um.proto.Message;
import com.ukelink.um.proto.Vmmbody;
import com.ukelink.um.proto.Vmmheader;
import okhttp3.MediaType;
import okhttp3.RequestBody;
import okio.BufferedSink;

import java.io.ByteArrayOutputStream;
import java.io.IOException;

import static java.nio.charset.StandardCharsets.UTF_8;

public class MySendRequestBody extends RequestBody {
    String messageId;
    public MySendRequestBody(String messageId){
        this.messageId=messageId;
        System.out.println("id_111: " + messageId);
    }
    Message.SendMessageRequest request;

    @Override
    public MediaType contentType() {
        return MediaType.parse("application/protobuf");
    }

    @Override
    public long contentLength() throws IOException {
        Clientinfo.ClientInfo.UserValue.Builder userValue = Clientinfo.ClientInfo.UserValue.newBuilder();
        //userValue.setId("11223344");
        userValue.setNumber("138");
        Clientinfo.ClientInfo.Builder fromClientInfo = Clientinfo.ClientInfo.newBuilder();
        fromClientInfo.setClientId("aaa");
        fromClientInfo.setDeviceType("PC");
        fromClientInfo.setUserType(Clientinfo.ClientInfo.UserType.sim);
        fromClientInfo.setUserValue(userValue);
        Vmmheader.VmmHeader.Builder vmmheader = Vmmheader.VmmHeader.newBuilder();

        vmmheader.setFromUser(fromClientInfo);

        Clientinfo.ClientInfo.UserValue.Builder toUerValue = Clientinfo.ClientInfo.UserValue.newBuilder();
        //toUerValue.setId("445566");
        toUerValue.setNumber("33");
        Clientinfo.ClientInfo.Builder toClientInfo = Clientinfo.ClientInfo.newBuilder();
        toClientInfo.setClientId("bbb");
        toClientInfo.setDeviceType("PC");
        toClientInfo.setUserType(Clientinfo.ClientInfo.UserType.sim);

        toClientInfo.setUserValue(toUerValue);
        vmmheader.addToUser(toClientInfo);
        vmmheader.setConversationType(1);
        vmmheader.setAction(1);
        vmmheader.setLanguage("zh");

        vmmheader.setOriginalMsgId(messageId);

        Vmmbody.VmmBody.Builder vmmbody = Vmmbody.VmmBody.newBuilder();
        vmmbody.setContentType(1);
        vmmbody.setMediaType(1);
//            String Content = IDUtil.getId();
        vmmbody.setContent(ByteString.copyFromUtf8("I am content"));

        Message.SendMessageRequest.Builder vmm = Message.SendMessageRequest.newBuilder();
        //vmm.getVmmHeader().toBuilder().setServerMsgId("ddddddddddddddddd");
        //vmm.setVmmHeader(vmmheader);
        //vmm.setVmmBody(vmmbody);
        request = vmm.build();
        return 4 + 8 + 12 + 8L + request.getSerializedSize();
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
}

