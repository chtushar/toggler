import { Tabs, TabsContent, TabsList, TabsTrigger } from '../../ui/tabs'
import General from './General'
import Team from './Team'

const Settings = () => {
  return (
    <div className="w-full h-full">
      <Tabs defaultValue="general">
        <TabsList>
          <TabsTrigger value="general">General</TabsTrigger>
          <TabsTrigger value="team">Team</TabsTrigger>
        </TabsList>
        <div className="mt-8">
          <TabsContent value="general">
            <General />
          </TabsContent>
          <TabsContent value="team">
            <Team />
          </TabsContent>
        </div>
      </Tabs>
    </div>
  )
}

export default Settings
