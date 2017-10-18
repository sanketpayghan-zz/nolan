
#include<Python.h>
#include "libcall.h"

static PyObject * call_parallel_api(PyObject *self, PyObject *args, PyObject *keywds) {
	const char *apiName, *apiData, *concurrentOn;
	char  *result;
	if (!PyArg_ParseTuple(args, "sss", &apiName, &apiData, &concurrentOn))
		return NULL;
	result = call_api(apiName, apiData, concurrentOn);
	printf("%s, %s, %s", apiName, apiData, concurrentOn);
	char * s = result;
	// long length = 0;
	// while(s){
	// 	length++;
	// 	s++;
	// }
	//return Py_BuildValue("y#", result, length);
	// printf("%s\n", result);
	return PyString_FromString(result);
}

static PyMethodDef NolanMethods[] = {
	{"call_parallel_api",  call_parallel_api, METH_VARARGS,
		"Call apis in parallel."},
	{NULL, NULL, 0, NULL}
};

PyMODINIT_FUNC
initnolan(void)
{
	(void) Py_InitModule("nolan", NolanMethods);
}
