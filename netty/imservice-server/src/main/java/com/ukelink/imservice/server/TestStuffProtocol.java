//package com.ukelink.imservice.server;
//
//import java.util.List;
//
//public class TestStuffProtocol {
//    private Integer id;
//    private String  name;
//    private String email;
//    private InnerStuff innerStuff;
//    List<InnerStuff> innerStuffs;
//
//    @Override
//    public String toString() {
//        return "TestStuffProtocol{" +
//                "id=" + id +
//                ", name='" + name + '\'' +
//                ", email='" + email + '\'' +
//                ", innerStuff=" + innerStuff +
//                ", innerStuffs=" + innerStuffs +
//                '}';
//    }
//
//    private enum InnerCarType
//    {
//        AUDI,
//        BENZ,
//        LAMBORGHINI,
//        DASAUTO;
//    }
//
//    public InnerStuff getInnerStuff() {
//        return innerStuff;
//    }
//
//    public void setInnerStuff(InnerStuff innerStuff) {
//        this.innerStuff = innerStuff;
//    }
//
//    public List<InnerStuff> getInnerStuffs() {
//        return innerStuffs;
//    }
//
//    public void setInnerStuffs(List<InnerStuff> innerStuffs) {
//        this.innerStuffs = innerStuffs;
//    }
//
//    public TestStuffProtocol(Integer id, String name, String email, InnerStuff innerStuff, List<InnerStuff> innerStuffs) {
//        this.id = id;
//        this.name = name;
//        this.email = email;
//        this.innerStuff = innerStuff;
//        this.innerStuffs = innerStuffs;
//    }
//
//    public Integer getId() {
//        return id;
//    }
//
//    public void setId(Integer id) {
//        this.id = id;
//    }
//
//    public String getName() {
//        return name;
//    }
//
//    public void setName(String name) {
//        this.name = name;
//    }
//
//    public String getEmail() {
//        return email;
//    }
//
//    public void setEmail(String email) {
//        this.email = email;
//    }
//
//    class InnerStuff{
//        private String name;
//        private InnerCarType innerCarType;
//
//        public InnerStuff(String name, InnerCarType innerCarType) {
//            this.name = name;
//            this.innerCarType = innerCarType;
//        }
//
//        public String getName() {
//            return name;
//        }
//
//        public void setName(String name) {
//            this.name = name;
//        }
//
//        public InnerCarType getInnerCarType() {
//            return innerCarType;
//        }
//
//        public void setInnerCarType(InnerCarType innerCarType) {
//            this.innerCarType = innerCarType;
//        }
//        @Override
//        public String toString() {
//            return "InnerStuff{" +
//                    "name='" + name + '\'' +
//                    ", innerCarType=" + innerCarType +
//                    '}';
//        }
//    }
//
//}
