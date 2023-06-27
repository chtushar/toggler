import type { SubmitHandler } from 'react-hook-form'
import * as z from 'zod'
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from '@/components/ui/form'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import useRegisterAdmin from '@/hooks/mutations/useRegisterAdmin'
import { Loader2 } from 'lucide-react'
import { useForm } from 'react-hook-form'
import { basicAuthSignUpSchema } from '@/lib/formValidators'
import { zodResolver } from '@hookform/resolvers/zod'

const RegisterAdmin = () => {
  const form = useForm<z.infer<typeof basicAuthSignUpSchema>>({
    resolver: zodResolver(basicAuthSignUpSchema),
    defaultValues: {
      name: '',
      email: '',
      password: '',
      confirmPassword: '',
    },
  })
  const { mutate, isLoading } = useRegisterAdmin()

  const handleSubmit: SubmitHandler<{
    name: string
    email: string
    password: string
    confirmPassword: string
  }> = data => {
    mutate({
      email: data.email,
      name: data.name,
      password: data.password,
    })
  }

  return (
    <div className="w-full h-full bg-background flex flex-col items-center justify-center">
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(handleSubmit)}
          className="w-full max-w-sm flex flex-col space-y-4"
        >
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Name</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
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
          <FormField
            control={form.control}
            name="confirmPassword"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Confirm Password</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit" disabled={isLoading}>
            {isLoading && <Loader2 className="w-4 h-4 mr-2 animate-spin" />}
            Register Admin
          </Button>
        </form>
      </Form>
    </div>
  )
}

export default RegisterAdmin
