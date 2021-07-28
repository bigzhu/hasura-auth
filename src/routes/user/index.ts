import { Router } from "express";
import { createValidator } from "express-joi-validation";

import { asyncWrapper as aw } from "@/helpers";
import {
  userActivateSchema,
  userEmailResetSchema,
  userEmailSchema,
  userMFASchema,
  userPasswordResetSchema,
  userPasswordSchema,
} from "@/validation";
import { userActivateHandler } from "./activate";
import { userMFAHandler } from "./mfa";
import { userHandler } from "./user";
import { userPasswordHandler } from "./password";
import { userPasswordResetHandler } from "./password-reset";
import { userEmailHandler } from "./email";
import { userEmailReset } from "./email/reset";

const router = Router();

router.get("/user", aw(userHandler));

router.post(
  "/user/password/reset",
  createValidator().body(userPasswordResetSchema),
  aw(userPasswordResetHandler)
);

router.post(
  "/user/password",
  createValidator().body(userPasswordSchema),
  aw(userPasswordHandler)
);

router.post(
  "/user/email/reset",
  createValidator().body(userEmailResetSchema),
  aw(userEmailReset)
);

router.post(
  "/user/email",
  createValidator().body(userEmailSchema),
  aw(userEmailHandler)
);

router.post(
  "/user/activate",
  createValidator().body(userActivateSchema),
  aw(userActivateHandler)
);

router.post(
  "/user/mfa",
  createValidator().body(userMFASchema),
  aw(userMFAHandler)
);

const userRouter = router;
export { userRouter };
