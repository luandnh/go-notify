package response

import (
	"net/http"
)

func Paging(data, prevPage, nextPage interface{}) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"code":      http.StatusOK,
		"message":   http.StatusText(http.StatusOK),
		"data":      data,
		"next_page": nextPage,
		"prev_page": prevPage,
	}
}

func Scroll(data interface{}, scrollId string) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"code":      http.StatusOK,
		"message":   http.StatusText(http.StatusOK),
		"items":     data,
		"scroll_id": scrollId,
	}
}

func Data(code int, data interface{}) (int, interface{}) {
	return code, map[string]interface{}{
		"code":    code,
		"message": http.StatusText(code),
		"data":    data,
	}
}

func OK(data interface{}) (int, interface{}) {
	return http.StatusOK, data
}

func Created(data map[string]interface{}) (int, interface{}) {
	result := map[string]interface{}{
		"created": true,
	}
	for key, value := range data {
		result[key] = value
	}
	return http.StatusCreated, result
}
func Error(code int, msg interface{}) (int, interface{}) {
	return code, map[string]interface{}{
		"code":    code,
		"message": http.StatusText(code),
		"error":   msg,
	}
}

func ServiceUnavailable() (int, interface{}) {
	return http.StatusServiceUnavailable, map[string]interface{}{
		"code":    http.StatusServiceUnavailable,
		"message": http.StatusText(http.StatusServiceUnavailable),
		"error":   http.StatusText(http.StatusServiceUnavailable),
	}
}

func ServiceUnavailableMsg(msg interface{}) (int, interface{}) {
	return http.StatusServiceUnavailable, map[string]interface{}{
		"code":    http.StatusServiceUnavailable,
		"message": http.StatusText(http.StatusServiceUnavailable),
		"error":   msg,
	}
}

func BadRequest() (int, interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": http.StatusText(http.StatusBadRequest),
		"error":   http.StatusText(http.StatusBadRequest),
	}
}

func BadRequestMsg(msg interface{}) (int, interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": http.StatusText(http.StatusBadRequest),
		"error":   msg,
	}
}

func NotFound() (int, interface{}) {
	return http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": http.StatusText(http.StatusNotFound),
		"error":   http.StatusText(http.StatusNotFound),
	}
}

func NotFoundMsg(msg interface{}) (int, interface{}) {
	return http.StatusNotFound, map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": http.StatusText(http.StatusNotFound),
		"error":   msg,
	}
}

func Forbidden() (int, interface{}) {
	return http.StatusForbidden, map[string]interface{}{
		"code":    http.StatusForbidden,
		"message": http.StatusText(http.StatusForbidden),
		"error":   "Do not have permission for the request.",
	}
}

func Unauthorized() (int, interface{}) {
	return http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"message": http.StatusText(http.StatusUnauthorized),
		"error":   http.StatusText(http.StatusUnauthorized),
	}
}

func EmptyData() []any {
	return []any{}
}

func Empty() any {
	return interface{}(nil)
}
