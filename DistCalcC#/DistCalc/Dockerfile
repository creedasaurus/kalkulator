FROM microsoft/aspnetcore:2.0-nanoserver-1709 AS base
WORKDIR /app
EXPOSE 80

FROM microsoft/aspnetcore-build:2.0-nanoserver-1709 AS build
WORKDIR /src
COPY DistCalc/DistCalc.csproj DistCalc/
RUN dotnet restore DistCalc/DistCalc.csproj
COPY . .
WORKDIR /src/DistCalc
RUN dotnet build DistCalc.csproj -c Release -o /app

FROM build AS publish
RUN dotnet publish DistCalc.csproj -c Release -o /app

FROM base AS final
WORKDIR /app
COPY --from=publish /app .
ENTRYPOINT ["dotnet", "DistCalc.dll"]
