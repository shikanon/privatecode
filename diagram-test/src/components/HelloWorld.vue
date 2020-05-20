<template>
  <div id="helloworld">
    <div id="mountNode"></div>
    <div id="test"></div>
  </div>
</template>

<script>
import G6 from '@antv/g6';



export default {
  name: 'HelloWorld',
  mounted() {
    this.initG6();
  },
  methods:{
    initG6(){
      const data = {
        // 点集
        nodes: [
          {
            id: 'pod1', // String，该节点存在则必须，节点的唯一标识
            shape: 'image', //节点类型
            img: 'https://raw.githubusercontent.com/mingrammer/diagrams/master/resources/k8s/compute/pod.png', //图片位置
            x: 100, // Number，可选，节点位置的 x 值
            y: 200, // Number，可选，节点位置的 y 值
            size: 100,
            label: 'Web APP',
            labelCfg: {
              position: 'bottom'
            }
          },
          {
            id: 'pod2', // String，该节点存在则必须，节点的唯一标识
            shape: 'image', //节点类型
            img: 'https://raw.githubusercontent.com/mingrammer/diagrams/master/resources/k8s/compute/pod.png', //图片位置
            x: 100, // Number，可选，节点位置的 x 值
            y: 400, // Number，可选，节点位置的 y 值
            size: 100,
            label: 'Web APP',
            labelCfg: {
              position: 'bottom'
            }
          },
          {
            id: 'pv2', // String，该节点存在则必须，节点的唯一标识
            shape: 'image', //节点类型
            img: 'https://raw.githubusercontent.com/mingrammer/diagrams/master/resources/k8s/storage/pv.png', //图片位置
            x: 300, // Number，可选，节点位置的 x 值
            y: 200, // Number，可选，节点位置的 y 值
            size: 100,
            label: 'Data',
            labelCfg: {
              position: 'bottom'
            }
          },
        ],
        // 边集
        edges: [
          {
            source: 'pod1', // String，必须，起始点 id
            target: 'pv2', // String，必须，目标点 id
          },
          {
            source: 'pod2', // String，必须，起始点 id
            target: 'pv2', // String，必须，目标点 id
          },
        ],
      };

      const graph = new G6.Graph({
        container: 'mountNode', // String | HTMLElement，必须，在 Step 1 中创建的容器 id 或容器本身
        width: 800, // Number，必须，图的宽度
        height: 500, // Number，必须，图的高度
        // 为方便演示，加粗边
        defaultEdge: {
          style: {
            lineWidth: 4,
          },
        },
      });

      graph.data(data); // 读取 Step 2 中的数据源到图上
      graph.render(); // 渲染图
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
