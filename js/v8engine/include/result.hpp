#include <string>

template <class T>
struct result
{
public:
    enum Status {
        Success,
        Error
    };

    // Feel free to change the default behavior... I use implicit
    // constructors for type T for syntactic sugar in return statements.
    result(T resultValue) : s(Success), v(resultValue) {}
    explicit result(Status status, std::string errMsg = std::string()) : s(status), v(), errMsg(errMsg) {}
    result() : s(Error), v() {} // Error without message
    result(T resultValue, std::string errMsg) : s(Error), v(resultValue), errMsg(errMsg) {}

    // Explicit error with message
    static result error(std::string errMsg) { return result(Error, errMsg); }

    // Implicit conversion to type T
    operator T() const { return v; }
    // Explicit conversion to type T
    T value() const { return v; }

    Status status() const { return s; }
    bool isError() const { return s == Error; }
    bool isSuccessful() const { return s == Success; }
    std::string errorMessage() const { return errMsg; }

private:
    T v;
    Status s;

    // if you want to provide error messages:
    std::string errMsg;
};
