package com.example.restapi.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.http.ResponseEntity;
import org.springframework.http.HttpStatus;

@RestController
public class UserController {

    @GetMapping("/user/{username}")
    public ResponseEntity<String> getUsername(@PathVariable String username) {
        return new ResponseEntity<>("{\"name\":\"" + username + "\"}", HttpStatus.OK);
    }
}