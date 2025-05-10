package by.andrewbesedin.distcomp.services;

import by.andrewbesedin.distcomp.dto.TweetMapper;
import by.andrewbesedin.distcomp.dto.TweetRequestTo;
import by.andrewbesedin.distcomp.dto.TweetResponseTo;
import by.andrewbesedin.distcomp.dto.*;
import by.andrewbesedin.distcomp.entities.Tag;
import by.andrewbesedin.distcomp.entities.Tweet;
import by.andrewbesedin.distcomp.repositories.TagRepository;
import by.andrewbesedin.distcomp.repositories.TweetRepository;
import lombok.AllArgsConstructor;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.cache.annotation.CacheEvict;
import org.springframework.cache.annotation.CachePut;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Collections;
import java.util.List;
import java.util.NoSuchElementException;
import java.util.Set;
import java.util.stream.Collectors;

@Service
@AllArgsConstructor
public class TweetService {
    public final TweetRepository repImpl;
    private final TagRepository tagRepository;
    @Qualifier("tweetMapper")
    public final TweetMapper mapper;

    public List<TweetResponseTo> getAll() {
        return repImpl.getAll().map(mapper::out).toList();
    }
    @Cacheable(value = "tweets", key = "#id")
    public TweetResponseTo getById(Long id) {
        return repImpl.get(id).map(mapper::out).orElseThrow();
    }
    @CachePut(value = "tweets", key = "#req.id")
    public TweetResponseTo create(TweetRequestTo req) {

        return repImpl.create(map(req)).map(mapper::out).orElseThrow();
    }
    private Tweet map(TweetRequestTo req) {
        Tweet tweet = mapper.in(req);
        Set<Tag> tags = (req.getTags() == null || req.getTags().isEmpty())
                ? Collections.emptySet()
                : req.getTags().stream()
                .map(name -> tagRepository.findByName(name).orElseGet(() -> {
                    Tag newTag = new Tag();
                    newTag.setName(name);
                    return tagRepository.save(newTag);
                }))
                .collect(Collectors.toSet());

        tweet.setTags(tags);
        return tweet;
    }
    @CachePut(value = "tweets", key = "#req.id")
    public TweetResponseTo update(TweetRequestTo req) {
        return repImpl.update(map(req)).map(mapper::out).orElseThrow();
    }
    @CacheEvict(value = "tweets", key = "#id")
    @Transactional
    public void delete(Long id) {
        Tweet tweet = repImpl.findById(id)
                .orElseThrow(() -> new NoSuchElementException("No element with id " + id));
        Set<Tag> tags = tweet.getTags();
        repImpl.delete(id);

        for (Tag tag : tags) {
            if (repImpl.countTweetsWithTag(tag.getId()) == 0) {
                tagRepository.delete(tag);
            }
        }

    }
}

