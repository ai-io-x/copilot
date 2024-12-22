package com.example.sqliteserver.controller;

import com.example.sqliteserver.model.SQLiteModel;
import com.example.sqliteserver.service.SQLiteService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/websites")
public class SQLiteController {

    @Autowired
    private SQLiteService sqliteService;

    @PostMapping
    public SQLiteModel addWebsite(@RequestBody SQLiteModel website) {
        return sqliteService.addWebsite(website);
    }

    @GetMapping
    public List<SQLiteModel> getAllWebsites() {
        return sqliteService.getAllWebsites();
    }

    @PutMapping("/{id}")
    public SQLiteModel updateWebsite(@PathVariable Long id, @RequestBody SQLiteModel website) {
        return sqliteService.updateWebsite(id, website);
    }

    @DeleteMapping("/{id}")
    public void deleteWebsite(@PathVariable Long id) {
        sqliteService.deleteWebsite(id);
    }
}