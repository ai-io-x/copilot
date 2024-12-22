package com.example.restserver.controller;

import com.example.restserver.model.DataModel;
import com.example.restserver.service.DataService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/api/data")
public class DataController {

    @Autowired
    private DataService dataService;

    @PostMapping
    public DataModel createData(@RequestBody DataModel data) {
        return dataService.createData(data);
    }

    @GetMapping
    public List<DataModel> getAllData() {
        return dataService.getAllData();
    }

    @GetMapping("/{id}")
    public Optional<DataModel> getDataById(@PathVariable Long id) {
        return dataService.getDataById(id);
    }

    @PutMapping("/{id}")
    public DataModel updateData(@PathVariable Long id, @RequestBody DataModel data) {
        return dataService.updateData(id, data);
    }

    @DeleteMapping("/{id}")
    public void deleteData(@PathVariable Long id) {
        dataService.deleteData(id);
    }
}