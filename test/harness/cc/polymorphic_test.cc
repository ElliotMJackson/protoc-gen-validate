#include "test/gen/package/bool.pb.h"
#include "test/harness/cc/other.pb.h"
#include "validate/validate.h"

int main() {
  test::harness::cc::Foo foo;

  // This does not have an associated validator but should still pass.
  std::string err;
  if (!pgv::BaseValidator::AbstractCheckMessage(foo, &err)) {
    return EXIT_FAILURE;
  }

  test::gen::package::BoolConstTrue bool_const_true;
  bool_const_true.set_val(false);
  if (pgv::BaseValidator::AbstractCheckMessage(bool_const_true, &err)) {
    return EXIT_FAILURE;
  }

  return EXIT_SUCCESS;
}
