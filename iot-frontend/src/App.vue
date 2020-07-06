<template>
	<div id="app">
		<el-container>
			<el-aside width="200px">
				<!-- <div href="L" class="logo" style="line-height: 3.125rem;">
					<img src="./assets/logo.png" style="width: 1.25rem;" />
				</div> -->
				<el-menu default-active="2" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" background-color="#33373d"
				 text-color="#fff" active-text-color="#ffd04b">
				 <el-menu-item  index="1" class="hui-iot-logo"><strong>HUI-IOT</strong></el-menu-item>
					<el-menu-item index="2">
						<template slot="title"><i class="el-icon-pie-chart"></i>数据统计</template>
					</el-menu-item>
					<el-submenu index="3">
						<template slot="title"><i class="el-icon-cpu"></i>设备管理</template>
						<el-menu-item-group>
							<el-menu-item index="3-1">类型管理</el-menu-item>
							<el-menu-item index="3-2">设备管理</el-menu-item>
						</el-menu-item-group>
					</el-submenu>
					<el-submenu index="4">
						<template slot="title"><i class="el-icon-s-data"></i>数据管理</template>
						<el-menu-item-group>
							<el-menu-item index="4-1">数据查询</el-menu-item>
							<el-menu-item index="4-2">事件定义</el-menu-item>
							<el-menu-item index="4-2">数据推送</el-menu-item>
						</el-menu-item-group>
					</el-submenu>
					<el-submenu index="5">
						<template slot="title"><i class="el-icon-setting"></i>系统管理</template>
						<el-menu-item-group>
							<el-menu-item index="5-1">租户管理</el-menu-item>
							<el-menu-item index="5-2">用户管理</el-menu-item>
							<el-menu-item index="5-2">角色管理</el-menu-item>
							<el-menu-item index="5-2">部门管理</el-menu-item>
						</el-menu-item-group>
					</el-submenu>
				</el-menu>
			</el-aside>
			<el-container>
				<el-header >
					<el-row>
						<el-col :span="24">
							<el-menu :default-active="activeIndex2" class="el-menu-demo" mode="horizontal" @select="handleSelect"
							 background-color="#33373d" text-color="#fff" active-text-color="#ffd04b">
								<el-menu-item style="float: right;" index="3"><a href="https://www.baidu.com">登出</a></el-menu-item>
								<el-submenu index="2" style="float: right;">
									<template slot="title">个人中心</template>
									<el-menu-item index="2-1">修改密码</el-menu-item>
									<el-menu-item index="2-2">用户信息</el-menu-item>
								</el-submenu>
							</el-menu>
						</el-col>
					</el-row>
				</el-header>
				<el-main></el-main>
			</el-container>
		</el-container>
	</div>
</template>

<script>
	export default {
		methods: {
			startHacking() {
				this.$notify({
					title: 'It works!',
					type: 'success',
					message: 'We\'ve laid the ground work for you. It\'s time for you to build something epic!',
					duration: 5000
				})
			}
		}
	}
</script>

<style>
	#app {
		font-family: Helvetica, sans-serif;
		text-align: center;
	}
	.hui-iot-logo {
		border-bottom: 1px;
		border-bottom-style: solid;
		border-bottom-color: #ffffff;
		text-align: center;
		height: 3.8125rem;
	}
	
	.el-header,
	.el-footer {
		padding: unset;
		background-color: #23262a;
		color: #FFFFFF;
		text-align: center;
	}

	.el-aside {
		border-right: solid 1px;
		border-right-color: #FFFFFF;
		background-color: #33373d;
		color: #333;
		text-align: left;
		height: 66.3125rem;
		/* padding-left: unset; */
		min-width:unset;
	}
	

	.el-main {
		background-color: #ecf1fa;
		color: #333;
		text-align: center;
		height: 62.5rem;
	}
	.el-menu{
		border-right: unset;
	}
	.el-submenu .el-menu-item{
		min-width: 199px;
		text-align: center;
		
	}
	.el-menu-vertical-demo{
		background-color: #161616;
		
	}
	.el-menu-item{
		background-color: #161616;
		
	}
	.el-container {
		margin-bottom: 40px;
	}

	.el-container:nth-child(5) .el-aside,
	.el-container:nth-child(6) .el-aside {
		line-height: 260px;
	}

	.el-container:nth-child(7) .el-aside {
		line-height: 320px;
	}
</style>

<script>
	export default {
		data() {
			return {
				activeIndex: '1',
				activeIndex2: '1'
			};
		},
		methods: {
			handleSelect(key, keyPath) {
				console.log(key, keyPath);
			}
		}
	}
</script>
