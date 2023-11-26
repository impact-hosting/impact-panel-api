from setuptools import setup


setup(
    name='impact_api',
    author='Impact Hosting',
    maintainer='DeathHound6',
    packages=[item for item in open('../requirements.txt').readlines() if item not in ['', '\n', '\r\n']],
    license='MIT',
    long_description=open('../README.md').read(),
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
