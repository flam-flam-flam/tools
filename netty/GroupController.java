package com.ukelink.um.group.controller;

import com.ukelink.um.group.bean.GroupIdAndCreateTime;
import com.ukelink.um.group.bean.VmpHeader;
import com.ukelink.um.group.service.GroupService;
import com.ukelink.um.group.utils.BaseResponseUtils;
import com.ukelink.um.group.utils.IdTools;
import com.ukelink.um.group.utils.ResponseCodeEnum;
import com.ukelink.um.proto.Baseresponse.BaseResponse;
import com.ukelink.um.proto.Group.CreateGroupRequest;
import com.ukelink.um.proto.Group.CreateGroupResponse;
import java.io.IOException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * 群组相关的控制层
 *
 * @author liyuan.liu
 * @date 2019-07-01 16:44
 */
@RestController
@Slf4j
public class GroupController {

    @Autowired
    private GroupService groupService;

    /**
     * 创建群组
     */
    @PostMapping("/um/group/createGroup")
    public void createGroup(HttpServletRequest request, HttpServletResponse response) throws IOException {
        String lid = IdTools.getId1();
        log.info("[lid:{}] receive POST request [/um/group/createGroup]", lid);
        // 处理vmpHeader和vmpBody，分开它们
        VmpHeader vmpHeader = VmpHeader.getVmpHeader(request.getInputStream());
        CreateGroupRequest createGroupRequest = CreateGroupRequest.parseFrom(request.getInputStream());
        log.info("[lid:{}] createGroup() requestBody: {}", lid, createGroupRequest);
        CreateGroupResponse createGroupResponse;
        try {
            // 创建群组，生成群组id
            GroupIdAndCreateTime groupIdAndCreateTime = groupService.createGroup(createGroupRequest, lid);
            // 构建返回对象
            BaseResponse baseResponse = BaseResponseUtils
                .getBaseResponse(ResponseCodeEnum.SUCCESS.getCode(), ResponseCodeEnum.SUCCESS.getMsg(),
                    createGroupRequest.getEventId(), lid);
            createGroupResponse = CreateGroupResponse.newBuilder().setBaseResponse(baseResponse)
                .setGroupId(groupIdAndCreateTime.getGroupId()).setCreateTime(groupIdAndCreateTime.getCreateTime())
                .build();
        } catch (Exception e) {
            log.error("[lid:{}]", lid, e);
            // 构建异常返回对象
            BaseResponse baseResponse = BaseResponseUtils
                .getBaseResponse(ResponseCodeEnum.ERROR.getCode(), ResponseCodeEnum.ERROR.getMsg(),
                    createGroupRequest.getEventId(), lid);
            createGroupResponse = CreateGroupResponse.newBuilder().setBaseResponse(baseResponse).build();
        }
        // 将返回内容写到输出流中
        VmpHeader.setVmpHeaderAndBody(response.getOutputStream(), vmpHeader, createGroupResponse.toByteArray());
        log.info("[lid:{}] createGroup() responseBody: {}", lid, createGroupResponse);
    }

}
