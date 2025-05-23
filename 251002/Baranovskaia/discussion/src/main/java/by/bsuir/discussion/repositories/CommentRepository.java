package by.bsuir.discussion.repositories;

import by.bsuir.discussion.domain.Comment;
import org.springframework.data.cassandra.repository.AllowFiltering;
import org.springframework.data.cassandra.repository.MapIdCassandraRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface CommentRepository extends MapIdCassandraRepository<Comment> {
    @AllowFiltering
    void deleteCommentByTopicIdAndId(Long id, Long uuid);
    @AllowFiltering //Bad solution, BUT generally we need to search all comments of certain topic, so
                    //this search is redundant and that's why topicId is a partition key
    Optional<Comment> findCommentById(Long id);
    Optional<Comment> findCommentByTopicIdAndId(Long id, Long uuid);
}