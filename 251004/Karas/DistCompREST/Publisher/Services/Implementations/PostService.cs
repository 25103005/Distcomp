﻿using AutoMapper;
using FluentValidation;
using Publisher.DTO.RequestDTO;
using Publisher.DTO.ResponseDTO;
using Publisher.Exceptions;
using Publisher.Infrastructure.Validators;
using Publisher.Models;
using Publisher.Repositories.Interfaces;
using Publisher.Services.Interfaces;

namespace Publisher.Services.Implementations;

public class PostService : IPostService
{
    private readonly IPostRepository _postRepository;
    private readonly IMapper _mapper;
    private readonly PostRequestDTOValidator _validator;
    
    public PostService(IPostRepository postRepository, 
        IMapper mapper, PostRequestDTOValidator validator)
    {
        _postRepository = postRepository;
        _mapper = mapper;
        _validator = validator;
    }
    
    public async Task<IEnumerable<PostResponseDTO>> GetPostsAsync()
    {
        var posts = await _postRepository.GetAllAsync();
        return _mapper.Map<IEnumerable<PostResponseDTO>>(posts);
    }

    public async Task<PostResponseDTO> GetPostByIdAsync(long id)
    {
        var post = await _postRepository.GetByIdAsync(id)
                      ?? throw new NotFoundException(ErrorCodes.PostNotFound, ErrorMessages.PostNotFoundMessage(id));
        return _mapper.Map<PostResponseDTO>(post);
    }

    public async Task<PostResponseDTO> CreatePostAsync(PostRequestDTO post)
    {
        await _validator.ValidateAndThrowAsync(post);
        var postToCreate = _mapper.Map<Post>(post);
        var createdPost = await _postRepository.CreateAsync(postToCreate);
        return _mapper.Map<PostResponseDTO>(createdPost);
    }

    public async Task<PostResponseDTO> UpdatePostAsync(PostRequestDTO post)
    {
        await _validator.ValidateAndThrowAsync(post);
        var postToUpdate = _mapper.Map<Post>(post);
        var updatedPost = await _postRepository.UpdateAsync(postToUpdate)
                             ?? throw new NotFoundException(ErrorCodes.PostNotFound, ErrorMessages.PostNotFoundMessage(post.Id));
        return _mapper.Map<PostResponseDTO>(updatedPost);
    }

    public async Task DeletePostAsync(long id)
    {
        if (!await _postRepository.DeleteAsync(id))
        {
            throw new NotFoundException(ErrorCodes.PostNotFound, ErrorMessages.PostNotFoundMessage(id));
        }
    }
}
