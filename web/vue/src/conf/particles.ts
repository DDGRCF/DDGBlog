// 背景动画设置

const ParticlesConf = {
  conf: {
    background: {
      color: {
        value: "#f1f2f6",
      },
    },
    fpsLimit: 60,
    interactivity: {
      detectsOn: "window",
      events: {
        onClick: {
          enable: true,
          mode: "push",
        },
        onHover: {
          enable: true,
          mode: "grab",
        },
        resize: true,
      },
      modes: {
        bubble: {
          distance: 200,
          duration: 2,
          opacity: 0.6,
          size: 40,
          speed: 2,
        },
        push: {
          quantity: 4,
        },
        grab: {
          distance: 100,
          duration: 0.6,
        },
      },
    },
    particles: {
      color: {
        value: "#686de0",
      },
      links: {
        color: "#95afc0",
        distance: 150,
        enable: true,
        opacity: 0.5,
        width: 1,
      },
      collisions: {
        enable: false,
      },
      move: {
        direction: "none",
        enable: true,
        outMode: "bounce",
        random: false,
        speed: 2,
        straight: false,
      },
      number: {
        density: {
          enable: true,
          value_area: 600,
        },
        value: 50,
      },
      opacity: {
        value: 0.5,
      },
      shape: {
        type: "circle",
      },
      size: {
        random: true,
        value: 5,
      },
    },
    detectRetina: true,
  },
};
export default ParticlesConf;
