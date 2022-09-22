package build.buf.pgv;

import build.buf.pvg.cases.Enum;
import org.junit.Test;

import static org.assertj.core.api.Assertions.assertThatThrownBy;

public class RequiredValidationTest {
    @Test
    public void requiredWorks() throws ValidationException {
        // Present
        RequiredValidation.required("x", Enum.Outer.getDefaultInstance());
        // Absent
        assertThatThrownBy(() -> RequiredValidation.required("x", null)).isInstanceOf(ValidationException.class);
    }
}
