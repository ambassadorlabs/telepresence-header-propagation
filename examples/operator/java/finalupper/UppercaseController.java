package finalupper;

import java.util.Locale;
import java.util.Optional;
import java.util.concurrent.ThreadLocalRandom;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class UppercaseController {
  private static final Logger logger = LoggerFactory.getLogger(UppercaseController.class);

  @GetMapping("/finalupper")
  public String index(@RequestParam("subject") Optional<String> subject) {
    if (subject.isPresent()) {
      System.out.println(subject.get().toUpperCase());
      return subject.get().toUpperCase();
    } else {
      return "UPPERCASE FAILURE";
    }
  }
}