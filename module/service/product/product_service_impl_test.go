package product

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	ProductModel "github.com/arudji1211/go-dts-07/module/model/product"
	ProductRepoMock "github.com/arudji1211/go-dts-07/module/repository/product/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ctx := GetTestGinContext()
	type (
		input struct {
			ctx *gin.Context
		}
		want struct {
			Product []ProductModel.Product
			err     error
		}
	)

	testCases := []struct {
		desc   string
		doMock func(repoMock *ProductRepoMock.MockProductRepository)
		input  input
		want   want
	}{
		{
			desc: "test service get all product not found",
			input: input{
				ctx: ctx,
			},
			want: want{
				Product: []ProductModel.Product{},
				err:     errors.New("Data Product Not Found"),
			},
			doMock: func(repoMock *ProductRepoMock.MockProductRepository) {
				repoMock.
					EXPECT().
					GetAll(ctx).
					Return([]ProductModel.Product{}, errors.New("Data Product Not Found")).
					MaxTimes(1).
					AnyTimes()
			},
		}, {
			desc: "test service get all product found",
			input: input{
				ctx: ctx,
			},
			want: want{
				Product: []ProductModel.Product{
					{
						Id:          1,
						UserId:      1,
						Title:       "Bantal",
						Description: "Bantal Guling",
					}, {
						Id:          2,
						UserId:      1,
						Title:       "Meja",
						Description: "Meja belajar",
					},
				},
				err: nil,
			},
			doMock: func(repoMock *ProductRepoMock.MockProductRepository) {
				repoMock.
					EXPECT().
					GetAll(ctx).
					Return([]ProductModel.Product{
						{
							Id:          1,
							UserId:      1,
							Title:       "Bantal",
							Description: "Bantal Guling",
						}, {
							Id:          2,
							UserId:      1,
							Title:       "Meja",
							Description: "Meja belajar",
						},
					}, nil).
					MaxTimes(1).
					AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			//by default
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := ProductRepoMock.NewMockProductRepository(ctrl)
			tC.doMock(repoMock)

			svc := ProductServiceImpl{
				ProductRepo: repoMock,
			}

			data, err := svc.GetAll(tC.input.ctx, 1, "admin")
			if tC.want.err != nil {
				assert.EqualError(t, err, tC.want.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tC.want.Product, data)
		})
	}
}

func TestGetById(t *testing.T) {
	ctx := GetTestGinContext()
	type (
		input struct {
			idProduct int
		}
		want struct {
			Product ProductModel.Product
			err     error
		}
	)

	testCases := []struct {
		desc   string
		doMock func(repoMock *ProductRepoMock.MockProductRepository)
		input  input
		want   want
	}{
		{
			desc: "test service get product by id found",
			input: input{
				idProduct: 1,
			},
			want: want{
				Product: ProductModel.Product{
					Id:          1,
					UserId:      1,
					Title:       "Bantal",
					Description: "Bantal Guling"},
				err: nil,
			},
			doMock: func(repoMock *ProductRepoMock.MockProductRepository) {
				repoMock.
					EXPECT().
					GetById(ctx, 1).
					Return(ProductModel.Product{
						Id:          1,
						UserId:      1,
						Title:       "Bantal",
						Description: "Bantal Guling",
					}, nil).
					MaxTimes(1).
					AnyTimes()
			},
		}, {
			desc: "test service get product by id not found",
			input: input{
				idProduct: 3,
			},
			want: want{
				Product: ProductModel.Product{},
				err:     errors.New("Data is Not Found"),
			},
			doMock: func(repoMock *ProductRepoMock.MockProductRepository) {
				repoMock.
					EXPECT().
					GetById(ctx, 3).
					Return(ProductModel.Product{}, errors.New("Data is Not Found")).
					MaxTimes(1).
					AnyTimes()
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			//by default
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoMock := ProductRepoMock.NewMockProductRepository(ctrl)
			tC.doMock(repoMock)

			svc := ProductServiceImpl{
				ProductRepo: repoMock,
			}

			data, err := svc.GetById(ctx, tC.input.idProduct)
			if tC.want.err != nil {
				assert.EqualError(t, err, tC.want.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tC.want.Product, data)
		})
	}
}

func GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}
