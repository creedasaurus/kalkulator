FROM microsoft/dotnet:2.1-sdk-alpine as builder  
 
RUN mkdir -p /root/src/app/subtract-service
WORKDIR /root/src/app/subtract-service
 
COPY SubtractServiceDotnetCore.csproj . 
RUN dotnet restore ./SubtractServiceDotnetCore.csproj 

COPY . .
RUN dotnet publish -c release -o published 

FROM microsoft/dotnet:2.1-sdk-alpine

WORKDIR /root/
COPY --from=builder /root/src/app/subtract-service/published .
ENV ASPNETCORE_URLS=http://+:5000
EXPOSE 5000
CMD ["dotnet", "./SubtractServiceDotnetCore.dll"]
