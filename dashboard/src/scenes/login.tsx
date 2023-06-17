import { useRef } from 'react'
import type { FormEvent } from 'react'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Loader2 } from 'lucide-react'
import useLogin from '@/hooks/mutations/useLogin'
import { Button } from '@/components/ui/button'
import { errors } from '@/constants/errors'

const Login = () => {
  const { isLoading, mutate, isError } = useLogin()
  const formRef = useRef<HTMLFormElement>(null)

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)
    const data = {
      email: formData.get('email') as string,
      password: formData.get('password') as string,
    }
    mutate(data, {
      onError: () => {
        formRef.current?.reset()
      },
    })
  }

  return (
    <div className="w-full h-full bg-background flex flex-col items-center justify-center">
      <form
        ref={formRef}
        className="w-full max-w-sm flex flex-col space-y-4"
        onSubmit={handleSubmit}
      >
        <Label htmlFor="email" className="space-y-2">
          <span>Email</span>
          <Input
            type="email"
            name="email"
            required
            aria-required={true}
            autoComplete="off"
          />
        </Label>
        <Label htmlFor="password" className="space-y-2">
          <span>Password</span>
          <Input
            type="password"
            name="password"
            required
            aria-required={true}
            autoComplete="off"
          />
        </Label>
        <Button>
          {isLoading ? <Loader2 className="w-4 h-4 mr-2 animate-spin" /> : null}
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
    </div>
  )
}

export default Login
