export enum UserRole {
  Admin = 'admin',
  Member = 'member',
}

export interface User {
  id: string
  name: string
  uuid: string
  created_at: string
  updated_at: string
  role: UserRole
  email: string
  email_verified: boolean
}

export interface Organization {
  id: string
  name: string
  uuid: string
  created_at: string
  updated_at: string
}

export interface Project {
  id: string
  name: string
  uuid: string
  org_id: string
  owner_id: string
  created_at: string
  updated_at: string
}

export interface Environment {
  id: string
  name: string
  api_keys: Array<string>
  uuid: string
  created_at: string
  updated_at: string
}

export interface FeatureFlag {
  id: string
  name: string
  uuid: string
  flag_type: FeatureFlagType
  enabled: boolean
  value: object
  updated_at: string
}

export enum FeatureFlagType {
  Boolean = 'boolean',
}
