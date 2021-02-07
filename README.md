![Go](https://github.com/mchirico/mock-playground/workflows/Go/badge.svg)


# Ref:
https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f

https://medium.com/better-programming/a-gomock-quick-start-guide-71bee4b3a6f1

# mock-playground
```bash
GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.4
go doc github.com/golang/mock/gomock
```

```bash
# See gogenerate in files..

mockgen -destination ../../mocks/mock_car.go -package=Mocks github.com/mchirico/mock-playground/pkg/car/imp CAR
mockgen -destination ../../mocks/mock_idea.go -package=Mocks github.com/mchirico/mock-playground/pkg/idea/imp IDEA


```



