using Application.Extensions;
using Asp.Versioning;
using Asp.Versioning.Builder;
using Infrastructure.Extensions;

var builder = WebApplication.CreateBuilder(args);

builder.Logging.ClearProviders();
builder.Logging.AddConsole();
builder
    .Services.AddApplicationServices()
    .AddMongoInfrastructure(builder.Configuration)
    .AddRepositories()
    .AddKafkaPublishers(builder.Configuration)
    .AddKafkaConsumers(builder.Configuration);

builder.Services.AddControllers();

builder
    .Services.AddApiVersioning(options =>
    {
        options.DefaultApiVersion = new ApiVersion(1, 0);
        options.ReportApiVersions = true;
        options.ApiVersionReader = new UrlSegmentApiVersionReader();
    })
    .AddMvc();

var app = builder.Build();

ApiVersionSet apiVersionSet = app.NewApiVersionSet()
    .HasApiVersion(new ApiVersion(1, 0))
    // .HasApiVersion(new ApiVersion(2,0))
    .ReportApiVersions()
    .Build();

RouteGroupBuilder versionedGroup = app.MapGroup("api/v{version:apiVersion}")
    .WithApiVersionSet(apiVersionSet);

versionedGroup.MapControllers();

app.Run();
