import { Tabs, TabsList, TabsTrigger } from '@/components/ui/tabs'

const Organization = () => {
  return (
    <div className="p-4">
      <Tabs defaultValue="projects">
        <TabsList>
          <TabsTrigger value="projects">Projects</TabsTrigger>
          <TabsTrigger value="members">Team</TabsTrigger>
          <TabsTrigger value="settings">Settings</TabsTrigger>
        </TabsList>
      </Tabs>
    </div>
  )
}

export default Organization
