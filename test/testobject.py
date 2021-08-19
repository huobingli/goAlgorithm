#经典类或者旧试类

class A:
    pass

a = A()
#print A.__class__
print (a.__class__)
print (type(A))
print (type(a))


#新式类

class B(object):
    pass


b = B()

print (B.__class__)
print (b.__class__)
print (type(B))
print (type(b))
