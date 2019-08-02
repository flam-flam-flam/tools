package com.ukelink.um.dao;

import com.ukelink.um.entity.TableA;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.List;

/**
 * MapperTableA继承基类
 */
@Repository
public interface MapperTableA {

    int insert(TableA record);

    int insertSelective(TableA record);
    List<TableA> select();
}