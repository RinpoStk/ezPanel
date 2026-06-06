import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Server } from 'lucide-react'

export default function Services() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-bold text-foreground">服务管理</h2>
        <Button size="sm">
          <Server className="mr-2 h-4 w-4" />
          添加服务
        </Button>
      </div>
      <Card>
        <CardHeader>
          <CardTitle>服务列表</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex items-center justify-center py-12 text-muted-foreground">
            <p>暂无服务，点击"添加服务"开始</p>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
