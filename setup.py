from distutils.core import setup, Extension
import os

path_str = os.getcwd()
module1 = Extension('nolan',
					libraries = ['call'],
					library_dirs = [path_str],
					sources = ['nolan.c'])
setup (name = 'PackageName',version = '1.0',description = 'This is module can be used to make paralle http requests.',ext_modules = [module1])
