
#include<Python.h>
#include "libcall.h"

const CONNECTION_LIMIT = 35;

static PyObject * call_parallel_api(PyObject *self, PyObject *args, PyObject *keywds) {
	//const char *apiName, *apiData, *concurrentOn;
	char *apiName, *apiData, *concurrentOn;
	char  *result;
	if (!PyArg_ParseTuple(args, "sss", &apiName, &apiData, &concurrentOn))
		return NULL;
	result = call_api(apiName, apiData, concurrentOn);
	printf("%s, %s, %s", apiName, apiData, concurrentOn);
	//char * s = result;
	// long length = 0;
	// while(s){
	// 	length++;
	// 	s++;
	// }
	//return Py_BuildValue("y#", result, length);
	// printf("%s\n", result);
	return PyString_FromString(result);
}


static PyObject * call_parallel_rpc(PyObject *self, PyObject *args, PyObject *keywds) {
	//const char *apiName, *apiData, *concurrentOn;
	char *apiName, *apiData, *concurrentOn;
	char  *result;
	if (!PyArg_ParseTuple(args, "sss", &apiName, &apiData, &concurrentOn))
		return NULL;
	result = call_rpc(apiName, apiData, concurrentOn);
	printf("%s, %s, %s", apiName, apiData, concurrentOn);
	//char * s = result;
	// long length = 0;
	// while(s){
	// 	length++;
	// 	s++;
	// }
	//return Py_BuildValue("y#", result, length);
	// printf("%s\n", result);
	return PyString_FromString(result);
}

static PyObject * parallel_rpc_with_data(PyObject *self, PyObject *args, PyObject *keywds) {
	//const char *apiName, *apiData, *concurrentOn;
	char *apiName, *apiData;
	char  *result;
	if (!PyArg_ParseTuple(args, "ss", &apiName, &apiData))
		return NULL;
	result = call_rpc_with_data(apiName, apiData);
	//printf("%s, %s", apiName, apiData);
	//char * s = result;
	// long length = 0;
	// while(s){
	// 	length++;
	// 	s++;
	// }
	//return Py_BuildValue("y#", result, length);
	// printf("%s\n", result);
	return PyString_FromString(result);
}

static PyObject * post(PyObject *self, PyObject *args, PyObject *keywds) {
	char *apiName, *apiData;
	char  *result;
	if (!PyArg_ParseTuple(args, "ss", &apiName, &apiData))
		return NULL;
	result = postCall(apiName, apiData);
	return PyString_FromString(result);
}

static PyMethodDef NolanMethods[] = {
	{"call_parallel_api",  call_parallel_api, METH_VARARGS,
		"Call apis in parallel."},
	{"call_parallel_rpc", call_parallel_rpc, METH_VARARGS,
		"Make rpc requests parallel with same data."},
	{"parallel_rpc_with_data", parallel_rpc_with_data, METH_VARARGS,
		"Make parallel rpc requests with data."},
	{"post", post, METH_VARARGS,
		"Make parallel post requests with data."},
	{NULL, NULL, 0, NULL}
};

PyMODINIT_FUNC
initnolan(void)
{
	(void) Py_InitModule("nolan", NolanMethods);
}
