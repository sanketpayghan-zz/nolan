from distutils.core import setup, Extension

module1 = Extension('nolan',
					libraries = ['call'],
					library_dirs = ['/home/Happay_v2/Happay_v2/nolan'],
					sources = ['nolan.c'])
setup (name = 'PackageName',version = '1.0',description = 'This is module can be used to make paralle http requests.',ext_modules = [module1])
