from setuptools import setup, find_packages

bad_chars = ['\n', '\r\n']
items = [item for item in open('./requirements.txt').readlines() if item not in bad_chars]
packages = []
for item in items:
    for char in bad_chars:
        if item.endswith(char):
            item = item.replace(char, '')
    packages.append(item)

with open('../README.md') as readme:
    long_description = readme.read()

setup(
    name='impact_api',
    author='Impact Hosting',
    maintainer='DeathHound6',
    platforms=['unix', 'linux'],
    packages=find_packages(include=packages),
    license='MIT',
    long_description=long_description,
    python_requires='>=3.8',
    classifiers=[
        'Development Status :: 5 - Production/Stable',
        'Environment :: Web Environment',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Operating System :: POSIX :: Linux',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.9',
        'Programming Language :: Python :: 3.10',
        'Programming Language :: Python :: 3.11',
        'Topic :: Internet :: WWW/HTTP :: WSGI'
    ]
)
