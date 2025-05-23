package com.example.discussion;

import org.springframework.data.cassandra.repository.CassandraRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface RepositoryC extends CassandraRepository<CommentEntity, Long> {
}
