package com.example.sqliteserver.repository;

import com.example.sqliteserver.model.SQLiteModel;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface SQLiteRepository extends JpaRepository<SQLiteModel, Long> {
}