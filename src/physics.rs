use std::ops::{Add, AddAssign, Mul};

#[derive(Debug, Clone)]
pub struct Point(f64, f64);
impl Point {
    pub fn new(x: f64, y: f64) -> Self {
        return Self(x, y);
    }
}

impl Add<Point> for Point {
    type Output = Point;

    fn add(self, Point(x, y): Point) -> Self::Output {
        return Point(
            self.0 + x,
            self.1 + y
        );
    }
}

#[derive(Debug, Clone)]
pub struct Vector2D(f64, f64);

impl Vector2D {
    pub fn new(x: f64, y: f64) -> Vector2D {
        return Vector2D(x, y);
    }
}

impl AddAssign<Vector2D> for Vector2D {
    fn add_assign(&mut self, Vector2D(x, y): Vector2D) {
        self.0 += x;
        self.1 += y;
    }
}

impl Mul<f64> for Vector2D {
    type Output = Vector2D;

    fn mul(self, rhs: f64) -> Self::Output {
        return Vector2D(
            self.0 * rhs,
            self.1 * rhs,
        );
    }
}
