import { useRef } from 'react'
import type { SubmitHandler } from 'react-hook-form'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import * as z from 'zod'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Loader2 } from 'lucide-react'
import useLogin from '@/hooks/mutations/useLogin'
import { Button } from '@/components/ui/button'
import { errors } from '@/constants/errors'
import { basicAuthLoginSchema } from '@/lib/formValidators'

const Login = () => {
  const { isLoading, mutate, isError } = useLogin()
  const formRef = useRef<HTMLFormElement>(null)
  const form = useForm<z.infer<typeof basicAuthLoginSchema>>({
    resolver: zodResolver(basicAuthLoginSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  })

  const handleSubmit: SubmitHandler<{
    email: string
    password: string
  }> = data => {
    mutate(data, {
      onError: () => {
        formRef.current?.reset()
      },
    })
  }

  return (
    <div className="w-full h-full bg-background flex flex-col items-center justify-center">
      <Form {...form}>
        <form
          ref={formRef}
          className="w-full max-w-sm flex flex-col space-y-4"
          onSubmit={form.handleSubmit(handleSubmit)}
        >
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input {...field} type="password" />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button>
            {isLoading ? (
              <Loader2 className="w-4 h-4 mr-2 animate-spin" />
            ) : null}
            Login
          </Button>
          <div className="h-10">
            {isError ? (
              <p className="text-red-500 text-sm text-center">
                {errors['incorrect_credentials']}
              </p>
            ) : null}
          </div>
        </form>
      </Form>
    </div>
  )
}

export default Login
