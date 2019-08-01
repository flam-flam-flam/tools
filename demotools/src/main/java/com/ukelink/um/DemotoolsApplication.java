package com.ukelink.um;

import com.ukelink.um.service.NettyService;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.cloud.openfeign.EnableFeignClients;

@EnableDiscoveryClient
@EnableFeignClients
@SpringBootApplication
public class DemotoolsApplication {
//	@Autowired
//	private NettyService nettyService;
	public static void main(String[] args) {
		SpringApplication.run(DemotoolsApplication.class, args);
		try {
			NettyService.init();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}
