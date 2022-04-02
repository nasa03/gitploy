import { useState } from "react"
import { shallowEqual } from "react-redux"
import { Menu, Row, Col, Button, Drawer, Avatar, Dropdown, Badge, Space} from "antd"
import { SettingFilled } from "@ant-design/icons"

import { useAppSelector } from "../../redux/hooks"

import RecentActivities from "../../components/RecentActivities"

import Logo from "../../logo.svg"

export default function Header(): JSX.Element {
    const { 
        authorized, 
        user,
        deployments,
        reviews,
    } = useAppSelector(state => state.main, shallowEqual)

    const activitiesCount = deployments.length + reviews.length

    const [ isRecentActivitiesVisible, setRecentActivitiesVisible ] = useState(false)

    return (
        <Row>
            <Col span="16">
                <Menu theme="dark" mode="horizontal" >
                    <Menu.Item key={0}>
                        <a href="/"><img src={Logo} style={{width: 62}}/></a>
                    </Menu.Item>
                    <Menu.Item key={1}>
                        <a href="/">Home</a>
                    </Menu.Item>
                    {(user?.admin)? 
                        <Menu.Item key={2}>
                            <a href="/members">Members</a>
                        </Menu.Item>
                        : 
                        <></>}
                </Menu>
            </Col>
            <Col span="8" style={{textAlign: "right"}}>
                <Space>
                <Badge 
                    size="small" 
                    count={activitiesCount}
                >
                    <Button 
                        type="primary" 
                        shape="circle" 
                        icon={<SettingFilled spin={activitiesCount !== 0 }/>}
                        onClick={() => setRecentActivitiesVisible(true)}
                    />
                </Badge>
                <Drawer title="Recent Activities"
                    placement="right"
                    width={400}
                    visible={isRecentActivitiesVisible}
                    onClose={() => setRecentActivitiesVisible(false)}
                >
                    <RecentActivities 
                        deployments={deployments}
                        reviews={reviews}
                    />
                </Drawer>
                {(authorized) ? 
                    <Dropdown 
                        overlay={(
                            <Menu>
                                <Menu.Item key="0">
                                    <a href="/settings">Settings</a>
                                </Menu.Item>
                                <Menu.Divider />
                                <Menu.Item key="1">
                                    <a href="https://www.gitploy.io/docs/" target="_blank">Read Docs</a>
                                </Menu.Item>
                                <Menu.Item key="2">
                                    <a href="/signout">Sign out</a>
                                </Menu.Item>
                            </Menu>
                        )}>
                        <Avatar  src={user?.avatar}/>
                    </ Dropdown>
                    :
                    <a href="/" style={{color: "white"}}>Sign in</a>}
                </Space>
            </Col>
        </Row>
    )
}