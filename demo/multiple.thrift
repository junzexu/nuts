// copy from https://thrift-tutorial.readthedocs.io/en/latest/usage-example.html
namespace java tutorial
namespace py tutorial
namespace go tutorial

typedef i32 int // We can use typedef to get pretty names for the types we are using
service MultiplicationService
{
    int multiply(1:int n1, 2:int n2),
}