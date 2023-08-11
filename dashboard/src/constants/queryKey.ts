export const queryKey = {
  user: () => ['USER'],

  // Organizations
  organization: (orgId: string) => ['ORGANIZATION', orgId],
  userOrganizations: (userId: string) => ['ORGANIZATIONS', userId],
  organizationMembers: (orgId: string) => ['MEMBERS', orgId],

  // Projects
  project: (orgId: string, userId: string) => ['PROJECT', orgId, userId],
  projects: (orgId: string) => ['PROJECTS', orgId],

  // Environments
  environments: (projectId: string) => ['ENVIRONMENTS', projectId],

  // Feature Flags
  featureFlags: (projectId: string, environmentId: string) => [
    'FEATURE_FLAGS',
    projectId,
    environmentId,
  ],
}
