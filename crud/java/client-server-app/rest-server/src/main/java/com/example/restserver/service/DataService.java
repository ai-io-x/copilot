package com.example.restserver.service;

import com.example.restserver.model.DataModel;
import com.example.restserver.repository.DataRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class DataService {

    @Autowired
    private DataRepository dataRepository;

    public DataModel createData(DataModel data) {
        return dataRepository.save(data);
    }

    public List<DataModel> getAllData() {
        return dataRepository.findAll();
    }

    public Optional<DataModel> getDataById(Long id) {
        return dataRepository.findById(id);
    }

    public DataModel updateData(Long id, DataModel data) {
        if (dataRepository.existsById(id)) {
            data.setId(id);
            return dataRepository.save(data);
        }
        return null;
    }

    public void deleteData(Long id) {
        dataRepository.deleteById(id);
    }
}