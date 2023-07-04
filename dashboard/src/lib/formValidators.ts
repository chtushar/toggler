import { FeatureFlagType } from '@/types/models'
import * as z from 'zod'

export const basicAuthSignUpSchema = z
  .object({
    name: z.string().min(3),
    email: z.string().email(),
    password: z.string().min(4),
    confirmPassword: z.string().min(4),
  })
  .required()
  .superRefine(({ confirmPassword, password }, ctx) => {
    if (confirmPassword !== password) {
      ctx.addIssue({
        code: 'custom',
        message: 'The passwords did not match',
      })
    }
  })

export const basicAuthLoginSchema = z
  .object({
    email: z.string().email(),
    password: z.string().min(4),
  })
  .required()

export const orgGeneralSettings = z
  .object({
    name: z.string(),
  })
  .required()

export const createFeatureFlagSchema = z.object({
  name: z.string().min(3),
  type: z.enum(['boolean']),
})
