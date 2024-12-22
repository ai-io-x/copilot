package com.example.sqliteserver.service;

import com.example.sqliteserver.model.SQLiteModel;
import com.example.sqliteserver.repository.SQLiteRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class SQLiteService {

    @Autowired
    private SQLiteRepository sqliteRepository;

    public SQLiteModel addWebsite(SQLiteModel website) {
        return sqliteRepository.save(website);
    }

    public List<SQLiteModel> getAllWebsites() {
        return sqliteRepository.findAll();
    }

    public SQLiteModel updateWebsite(Long id, SQLiteModel website) {
        if (sqliteRepository.existsById(id)) {
            website.setId(id);
            return sqliteRepository.save(website);
        }
        return null;
    }

    public void deleteWebsite(Long id) {
        sqliteRepository.deleteById(id);
    }
}